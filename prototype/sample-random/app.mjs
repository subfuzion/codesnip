import process from 'node:process';
import express from 'express';

const app = express();
const port = parseInt(process.env.PORT, 10) || 8080;

app.get('/', async (req, res) => {

  // Roll the dice!
  const side = Math.floor(Math.random() * 6 + 1)
  const message = `🎲  You rolled: ${side}`;
  return res.send(message);
});

app.listen(port, () => {
  console.log(`Listening on ${port}`);
});
