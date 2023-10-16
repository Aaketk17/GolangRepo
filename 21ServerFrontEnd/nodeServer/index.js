const express = require('express')
const app = express()
const port = 8000

app.use(express.json())
app.use(express.urlencoded({extended: true}))

app.get('/', (req, res) => {
  res.status(200).send('Welcome to Node server')
})

app.get('/get', (req, res) => {
  res.status(200).json('Hello from node server')
})

app.post('/post', (req, res) => {
  let myJson = req.body
  res.status(200).send(myJson)
})

app.post('/postform', (req, res) => {
  res.status(200).send(JSON.stringify(req.body))
})

app.listen(port, () => {
  console.log(`Example node server is listening at http://localhost:${port}`)
})
