import { useEffect, useState } from 'react'
import './App.css'

const serverUrl = 'http://localhost:8080'

function App() {
  const [message, setMessage] = useState('')

  useEffect(() => {
    fetch(`${serverUrl}/api`)
      .then((res) => res.json())
      .then((data) => {
        setMessage(data.message)
      })
  })

  return (
    <>
      <div>
        {message}
      </div>
    </>
  )
}

export default App
