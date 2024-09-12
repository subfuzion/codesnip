import process from 'node:process';
import express from 'express';

const app = express();
const port = parseInt(process.env.PORT, 10) || 8080;

// Send a default greeting.
app.get('/', (req, res) => {
  res.send(`Hello World!`);
});

// Send a custom greeting.
app.get('/greet/:name', (req, res) => {
  const name = req.params.name;
  res.send(`Hello ${name}!`);
});

// Return not found message on any other route.
app.use((req, res) => {
  res.status(404).send(`Not Found: ${req.path}`);
})

const server = app.listen(port, () => {
  // Handle shut down signal gracefully.
  process.once('SIGTERM', () => {
    server.close(err => {
      if (err) {
        console.error(`Error stopping helloworld on ${port}: ${err.message}`);
      } else {
        console.log(`helloworld: stopped on port ${port}`);
      }
    });
  });
  console.log(`helloworld: listening on ${port}`);
});

