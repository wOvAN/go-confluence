package confluence

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Content struct {
	Id     string `json:"id"`
	Type   string `json:"type"`
	Status string `json:"status"`
	Title  string `json:"title"`
	Body   struct {
		Storage struct {
			Value          string `json:"value"`
			Representation string `json:"representation"`
		} `json:"storage"`
	} `json:"body"`
	Version struct {
		Number int `json:"number"`
	} `json:"version"`
	Space struct {
		Key string `json:"key"`
	} `json:"space,omitempty"`
}

func (w *Wiki) contentEndpoint(contentID string) (*url.URL, error) {
	return url.ParseRequestURI(w.endPoint.String() + "/content/" + contentID)
}

func (w *Wiki) DeleteContent(contentID string) error {
	contentEndPoint, err := w.contentEndpoint(contentID)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("DELETE", contentEndPoint.String(), nil)
	if err != nil {
		return err
	}

	res, err := w.sendRequest(req)
	if err != nil {
		return err
	}
	if Debug {
		log.Println("DeleteContent: ", string(res))
	}
	return nil
}

func (w *Wiki) GetContent(contentID string, expand []string) (*Content, error) {
	contentEndPoint, err := w.contentEndpoint(contentID)
	if err != nil {
		return nil, err
	}
	data := url.Values{}
	data.Set("expand", strings.Join(expand, ","))
	contentEndPoint.RawQuery = data.Encode()

	req, err := http.NewRequest("GET", contentEndPoint.String(), nil)
	if err != nil {
		return nil, err
	}

	res, err := w.sendRequest(req)
	if err != nil {
		return nil, err
	}

	if Debug {
		log.Println("GetContent: ", string(res))
	}

	var content Content
	err = json.Unmarshal(res, &content)
	if err != nil {
		return nil, err
	}

	return &content, nil
}

func (w *Wiki) UpdateContent(content *Content) (*Content, error) {
	jsonbody, err := json.Marshal(content)
	if err != nil {
		return nil, err
	}

	contentEndPoint, err := w.contentEndpoint(content.Id)
	req, err := http.NewRequest("PUT", contentEndPoint.String(), strings.NewReader(string(jsonbody)))
	req.Header.Add("Content-Type", "application/json")

	res, err := w.sendRequest(req)
	if err != nil {
		return nil, err
	}

	if Debug {
		log.Println("UpdateContent: ", string(res))
	}

	var newContent Content
	err = json.Unmarshal(res, &newContent)
	if err != nil {
		return nil, err
	}

	return &newContent, nil
}
