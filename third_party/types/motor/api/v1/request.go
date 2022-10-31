package v1

import "errors"

func (r *GenerateBucketRequest) Validate() error {
	hasBucket := r.Bucket != nil
	hasCreator := r.Creator != ""
	hasName := r.Name != ""
	hasUuid := r.Uuid != ""

	if !hasBucket && !hasCreator && !hasName && !hasUuid {
		return errors.New("No bucket, creator, name, or uuid provided")
	}

	if hasBucket {
		return nil
	}

	if hasUuid {
		return nil
	}

	if hasCreator && hasName {
		return nil
	} else {
		return errors.New("Must provide creator and name if not providing bucket or uuid")
	}
}

func (r *AddBucketItemsRequest) Validate() error {
	hasBucket := r.Bucket != nil
	hasCreator := r.Creator != ""
	hasName := r.Name != ""
	hasUuid := r.Uuid != ""

	if !hasBucket && !hasCreator && !hasName && !hasUuid {
		return errors.New("No bucket, creator, name, or uuid provided")
	}

	if hasBucket {
		return nil
	}

	if hasUuid {
		return nil
	}

	if hasCreator && hasName {
		return nil
	} else {
		return errors.New("Must provide creator and name if not providing bucket or uuid")
	}
}

func (r *GetBucketItemsRequest) Validate() error {
	hasBucket := r.Bucket != nil
	hasCreator := r.Creator != ""
	hasName := r.Name != ""
	hasUuid := r.Uuid != ""

	if !hasBucket && !hasCreator && !hasName && !hasUuid {
		return errors.New("No bucket, creator, name, or uuid provided")
	}

	if hasBucket {
		return nil
	}

	if hasUuid {
		return nil
	}

	if hasCreator && hasName {
		return nil
	} else {
		return errors.New("Must provide creator and name if not providing bucket or uuid")
	}
}

func (r *BurnBucketRequest) Validate() error {
	hasBucket := r.Bucket != nil
	hasCreator := r.Creator != ""
	hasName := r.Name != ""
	hasUuid := r.Uuid != ""

	if !hasBucket && !hasCreator && !hasName && !hasUuid {
		return errors.New("No bucket, creator, name, or uuid provided")
	}

	if hasBucket {
		return nil
	}

	if hasUuid {
		return nil
	}

	if hasCreator && hasName {
		return nil
	} else {
		return errors.New("Must provide creator and name if not providing bucket or uuid")
	}
}

func (r *FindBucketConfigRequest) Validate() error {
	hasBucket := r.Bucket != nil
	hasCreator := r.Creator != ""
	hasName := r.Name != ""
	hasUuid := r.Uuid != ""

	if !hasBucket && !hasCreator && !hasName && !hasUuid {
		return errors.New("No bucket, creator, name, or uuid provided")
	}

	if hasBucket {
		return nil
	}

	if hasUuid {
		return nil
	}

	if hasCreator && hasName {
		return nil
	} else {
		return errors.New("Must provide creator and name if not providing bucket or uuid")
	}
}
