package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

// Item represents a single item in the XML.
type Item struct {
	ITEMID    string `xml:"ITEMID"`
	COLOR     string `xml:"COLOR"`
	ImageURL  string
	MINQTY    string `xml: "MINQTY"`
	QTYFILLED string `xml: "QTYFILLED"`
}

// Root represents the root of the XML structure.
type Root struct {
	Items []Item `xml:"ITEM"`
}

// ParseXMLData parses the XML string into a Root struct.
func ParseXMLData(xmlData string) ([]Item, error) {
	var root Root
	err := xml.Unmarshal([]byte(xmlData), &root)
	if err != nil {
		return nil, err
	}

	// Add ImageURL for each item.
	for i := range root.Items {
		root.Items[i].ImageURL = fmt.Sprintf("https://img.bricklink.com/ItemImage/PN/%s/%s.png", root.Items[i].COLOR, root.Items[i].ITEMID)
	}
	return root.Items, nil
}

func main() {
	// XML data as a string
	xmlData := `
<ROOT>
    <ITEM>
        <ITEMTYPE>P</ITEMTYPE>
        <ITEMID>3023</ITEMID>
        <COLOR>11</COLOR>
        <MAXPRICE>0.0000</MAXPRICE>
        <MINQTY>10</MINQTY>
        <QTYFILLED>10</QTYFILLED>
        <CONDITION>X</CONDITION>
        <REMARKS>Added from model loc-1902</REMARKS>
        <NOTIFY>N</NOTIFY>
    </ITEM>
    <ITEM>
        <ITEMTYPE>P</ITEMTYPE>
        <ITEMID>44728</ITEMID>
        <COLOR>86</COLOR>
        <MAXPRICE>0.0000</MAXPRICE>
        <MINQTY>6</MINQTY>
        <QTYFILLED>6</QTYFILLED>
        <CONDITION>X</CONDITION>
        <REMARKS>Added from model loc-1902</REMARKS>
        <NOTIFY>N</NOTIFY>
    </ITEM>
</ROOT>
`

	var Items Item

	xmlFile, err := os.Open("list.xml")
	if err != nil {
		fmt.Print(err)
		return
	}
	defer xmlFile.Close()

	byteValue, _ := io.ReadAll(xmlFile)

	xml.Unmarshal(byteValue, &Items)

	// Parse the XML data
	items, err := ParseXMLData(xmlFile)
	if err != nil {
		fmt.Println("Error parsing XML:", err)
		return
	}

	// WEB Stuff
	// HTML template
	var htmlTemplate = `
        <!DOCTYPE html>
        <html lang="en">
        <head>
            <meta charset="UTF-8">
            <meta name="viewport" content="width=device-width, initial-scale=1.0">
            <title>Whishlost Items</title>
            <style>
                body { font-family: Arial, sans-serif; margin: 20px; }
                .item { display: flex; align-items: center; margin-bottom: 20px; }
                .item img { margin-right: 15px; width: 50px; height: 50px; }
                .item div { line-height: 1.5; }
            </style>
        </head>
        <body>
            <h1>BrickLink Items</h1>
            {{range .}}
            <div class="item">
                <img src="{{.ImageURL}}" alt="Item Image">
                <div>
                    <strong>Item ID:</strong> {{.ITEMID}}<br>
                    <strong>Color:</strong> {{.COLOR}}<br>
                    <strong>Have:</storng> {{.QTYFILLED}}<br>
                    <strong>Need:</strong> {{.MINQTY}}
                </div>
            </div>
            {{end}}
        </body>
        </html>
     `
	// Create a handler for the "/" route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.New("webpage").Parse(htmlTemplate)
		if err != nil {
			http.Error(w, "Error loading template", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, items)
	})

	// Start the server
	fmt.Println("Server is running at http://0.0.0.0:5000/")
	http.ListenAndServe(":5000", nil)
}
