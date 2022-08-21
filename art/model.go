package art

type Art struct {
    ID uint
    Content string
}

// getArt get a piece of art by ID
func getArt(id uint) Art {
    // get from DB    
    return Art{}
}

func (a Art) Draw() string {
    return a.Content
}


