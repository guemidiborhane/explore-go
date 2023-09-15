package links

func (link *Link) JSONResponse() ResponseBody {
	return ResponseBody{
		ID:    link.ID,
		Link:  link.Link,
		Short: link.Short,
	}
}
