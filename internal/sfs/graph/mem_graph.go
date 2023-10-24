package graph

import (
	"context"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"

	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/graph"
	"github.com/cayleygraph/cayley/schema"
	"github.com/cayleygraph/quad"
	"github.com/cayleygraph/quad/voc"
)

type Person struct {
	// dummy field to enforce all object to have a <id> <rdf:type> <ex:Person> relation
	// means nothing for Go itself
	rdfType struct{} `quad:"@type > ex:Person"`
	ID      quad.IRI `json:"@id"`     // tag @id is a special one - graph node value will be stored in this field
	Name    string   `json:"ex:name"` // field name (predicate) may be written as json field name
	Age     int      `quad:"ex:age"`  // or in a quad tag
}

type Coords struct {
	// Object may be without id - it will be generated automatically.
	// It's also not necessary to have a type definition.
	Lat float64 `json:"ex:lat"`
	Lng float64 `json:"ex:lng"`
}


func NewGraph() error {
   // Define an "ex:" prefix for IRIs that will be expanded to "http://example.org".
	// "ex:name" will become "http://example.org/name"
	voc.RegisterPrefix("ex:", "http://example.org/")

	// Associate Go type with an IRI.
	// All Coords objects will now generate a <id> <rdf:type> <ex:Coords> triple.
	schema.RegisterType(quad.IRI("ex:Coords"), Coords{})

	sch := schema.NewConfig()
	// Override a function to generate IDs. Can be changed to generate UUIDs, for example.
	sch.GenerateID = func(_ interface{}) quad.Value {
		return quad.BNode(fmt.Sprintf("node%d", rand.Intn(1000)))
	}

	// File for your new BoltDB. Use path to regular file and not temporary in the real world
	tmpdir, err := ioutil.TempDir("", "example")
	if err != nil {
        return err
    }

	defer os.RemoveAll(tmpdir) // clean up

	// Initialize the database
	err = graph.InitQuadStore("bolt", tmpdir, nil)
	if err != nil {
        return err
    }

	// Open and use the database
	store, err := cayley.NewGraph("bolt", tmpdir, nil)
	if err != nil {
        return err
    }
	defer store.Close()
	qw := graph.NewWriter(store)

	// Save an object
	bob := Person{
		ID:   quad.IRI("ex:bob").Full().Short(),
		Name: "Bob", Age: 32,
	}
	fmt.Printf("saving: %+v\n", bob)
	id, err := sch.WriteAsQuads(qw, bob)
	if err != nil {
        return err
    }
	err = qw.Close()
	if err != nil {
        return err
    }

	fmt.Println("id for object:", id, "=", bob.ID) // should be equal

	// Get object by id
	var someone Person
	err = sch.LoadTo(nil, store, &someone, id)
	if err != nil {
        return err
    }
	fmt.Printf("loaded: %+v\n", someone)

	// Or get all objects of type Person
	var people []Person
	err = sch.LoadTo(nil, store, &people)
	if err != nil {
        return err
    }
	fmt.Printf("people: %+v\n", people)

	fmt.Println()

	// Store objects with no ID and type
	coords := []Coords{
		{Lat: 12.3, Lng: 34.5},
		{Lat: 39.7, Lng: 8.41},
	}
	qw = graph.NewWriter(store)
	for _, c := range coords {
		id, err = sch.WriteAsQuads(qw, c)
	if err != nil {
        return err
    }
		fmt.Println("generated id:", id)
	}
	err = qw.Close()
	if err != nil {
        return err
    }

	// Get coords back
	var newCoords []Coords
	err = sch.LoadTo(nil, store, &newCoords)
	if err != nil {
        return err
    }
	fmt.Printf("coords: %+v\n", newCoords)

	// Print quads
	fmt.Println("\nquads:")
	ctx := context.TODO()
	it := store.QuadsAllIterator()
	for it.Next(ctx) {
		fmt.Println(store.Quad(it.Result()))
	}
    if err := it.Err(); err != nil {
        return err
    }
    return nil
}
