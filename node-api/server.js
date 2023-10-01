require("dotenv").config();
const express = require("express");
const bodyParser = require("body-parser");
const app = express();
const db = require("./queries");

app.use(bodyParser.json());
app.use(
  bodyParser.urlencoded({
    extended: true,
  })
);

app.get("/", (request, response) => {
  response.json({ info: "Node.js, Express, and Postgres API" });
});

app.get("/posts", db.getALLPosts);
app.get("/posts/:id", db.getPostById);
app.post("/posts", db.createPost);
app.put("/posts/:id", db.updatePost);
app.delete("/posts/:id", db.deletePost);

app.listen(process.env.PORT, () => {
  console.log(`App running on port ${process.env.PORT}.`);
});
