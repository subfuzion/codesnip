import process from 'node:process';
import express from 'express';

const app = express();
const port = parseInt(process.env.PORT, 10) || 8080;

app.get('/', async (req, res) => {
  try {
    const jokeApi = 'https://official-joke-api.appspot.com/random_joke';
    const response = await fetch(jokeApi);

    // Joke API returned an error.
    if (!response.ok) {
      console.error(`Error (joke api): ${response.status}`);
      return res.status(404).send(`Can't fetch joke (try again later): ${response.status}`);
    }

    // Joke API returned a joke.
    const {setup, punchline} = await response.json();
    return res.send(`${setup}\n${punchline}`);

  } catch (error) {
    // Unexpected error sending the request.
    console.error(`Error (server): ${error}`);
    return res.status(500).send(`Server error`);
  }
});

app.listen(port, () => {
  console.log(`Listening on ${port}`);
});
