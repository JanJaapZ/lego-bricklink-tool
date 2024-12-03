from flask import Flask, render_template_string
import xml.etree.ElementTree as ET
import requests

# Flask application
app = Flask(__name__)

# XML data as a string
xml_data = """
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
"""

# Parse the XML
root = ET.fromstring(xml_data)
items = []

# Load items into a list
for item in root.findall("ITEM"):
    item_data = {
        "item_id": item.find("ITEMID").text,
        "color": item.find("COLOR").text,
        "image_url": f"https://img.bricklink.com/ItemImage/PN/{item.find('COLOR').text}/{item.find('ITEMID').text}.png"
    }
    items.append(item_data)

# HTML template for the Flask app
html_template = """
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>BrickLink Items</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 20px; }
        .item { display: flex; align-items: center; margin-bottom: 20px; }
        .item img { margin-right: 15px; width: 50px; height: 50px; }
        .item div { line-height: 1.5; }
    </style>
</head>
<body>
    <h1>BrickLink Items</h1>
    {% for item in items %}
    <div class="item">
        <img src="{{ item.image_url }}" alt="Item Image">
        <div>
            <strong>Item ID:</strong> {{ item.item_id }}<br>
            <strong>Color:</strong> {{ item.color }}
        </div>
    </div>
    {% endfor %}
</body>
</html>
"""

# Flask route to display the items
@app.route("/")
def display_items():
    return render_template_string(html_template, items=items)

# Run the Flask app
if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5000)
