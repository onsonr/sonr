package file

const K_BUF_CHUNK = 32000
const K_B64_CHUNK = 31998 // Adjusted for Base64 -- has to be divisible by 3
const K_QUEUE_SIZE = 16

type FileQueue struct {
	queue []*FileItem
}

// @ Adds Item to File Queue
func (fq *FileQueue) Enqueue(element *FileItem) {
	fq.queue = append(fq.queue, element) // Simply append to enqueue.
}

// @ Pops Item from File Queue
func (fq *FileQueue) Dequeue() *FileItem {
	file := fq.queue[0]     // The first element is the one to be dequeued.
	fq.queue = fq.queue[1:] // Slice off the element once it is dequeued.
	return file
}

// @ Returns Queue Length
func (fq *FileQueue) Count() int {
	return len(fq.queue)
}

// @ Checks if Queue does not have any elements
func (fq *FileQueue) IsEmpty() bool {
	return len(fq.queue) == 0
}

// @ Checks if Queue has any elements
func (fq *FileQueue) IsNotEmpty() bool {
	return len(fq.queue) > 0
}

// @ Helper: Chunks string based on B64ChunkSize ^ //
func chunkBase64(s string) []string {
	chunkSize := K_B64_CHUNK
	ss := make([]string, 0, len(s)/chunkSize+1)
	for len(s) > 0 {
		if len(s) < chunkSize {
			chunkSize = len(s)
		}
		// Create Current Chunk String
		ss, s = append(ss, s[:chunkSize]), s[chunkSize:]
	}
	return ss
}
