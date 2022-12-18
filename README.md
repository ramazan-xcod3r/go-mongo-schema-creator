<h1>Go MongoDB Schema Creator</h1>
<p>This project is a tool for generating MongoDB schemas from JSON data. It reads a JSON file and creates a BSON document that represents the data in the JSON file. The BSON document can be used to create a MongoDB collection with the correct structure for the data.</p>
<h2>Getting Started</h2>
<p>These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.</p>
<h3>Prerequisites</h3>
<p>You will need the following software installed on your machine:</p>
<ul>
<li>Go (https://golang.org/)</li>
<li>MongoDB (https://www.mongodb.com/)</li>
</ul>
<h3>Installing</h3>
<p>To install the project, follow these steps:</p>
<ol>
<li>Clone the repository: <code>git clone https://github.com/Ramazan-xcod3r/goMongoSchemaCreator.git</code></li>
<li>Change into the project directory: <code>cd goMongoSchemaCreator</code></li>
<li>Run the project: <code>go run main.go --test data.json</code></li>
</ol>
<p>The <code>--test</code> flag is optional and can be used to test the data insertion into the database. The <code>data.json</code> file is the JSON file that the tool will read and create the BSON document from.</p>
<h2>Usage</h2>
<p>To use the tool, run the following command:</p>
<pre><code>go run main.go [--test] [data.json]</code></pre>
<p>The <code>--test</code> flag is optional and can be used to test the data insertion into the database. The <code>data.json</code> file is the JSON file that the tool will read and create the BSON document from.</p>
<h2>Example</h2>

![image](https://user-images.githubusercontent.com/76232388/208276258-4aabb1a9-6fec-4a8c-a550-2cf142a24c28.png)

<h2>Authors</h2>
<ul>
<li>

**Ramazan-xcod3r** - [Github](https://github.com/Ramazan-xcod3r)

</li>
</ul>
<h2>License</h2>
<p>This project is licensed under the MIT License </p>
