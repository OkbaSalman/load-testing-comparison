const Pool = require("pg").Pool;
const pool = new Pool({
  user: process.env.USER,
  host: process.env.HOST,
  database: "postgres",
  password: process.env.PASSWORD,
  port: 5432,
});

const getALLPosts = (request, response) => {
  pool.query("SELECT * FROM posts ORDER BY id ASC", (error, results) => {
    if (error) {
      throw error;
    }
    response.status(200).json(results.rows);
  });
};

const getPostById = (request, response) => {
  const id = parseInt(request.params.id);

  pool.query("SELECT * FROM posts WHERE id = $1", [id], (error, results) => {
    if (error) {
      throw error;
    }
    response.status(200).json(results.rows);
  });
};

const createPost = (request, response) => {
  const { Title, Body } = request.body;

  pool.query(
    "INSERT INTO posts (Title, Body) VALUES ($1, $2)",
    [Title, Body],
    (error, results) => {
      if (error) {
        throw error;
      }
      response.status(201).send(`Post added with ID: ${results.insertId}`);
    }
  );
};

const updatePost = (request, response) => {
  const id = parseInt(request.params.id);
  const { Title, Body } = request.body;

  pool.query(
    "UPDATE posts SET Title = $1, Body = $2 WHERE id = $3",
    [Title, Body, id],
    (error, results) => {
      if (error) {
        throw error;
      }
      response.status(200).send(`Post modified with ID: ${id}`);
    }
  );
};

const deletePost = (request, response) => {
  const id = parseInt(request.params.id);

  pool.query("DELETE FROM posts WHERE id = $1", [id], (error, results) => {
    if (error) {
      throw error;
    }
    response.status(200).send(`Post deleted with ID: ${id}`);
  });
};

module.exports = {
  getALLPosts,
  getPostById,
  createPost,
  updatePost,
  deletePost,
};
