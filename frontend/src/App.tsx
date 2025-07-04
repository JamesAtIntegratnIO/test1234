import { useState } from 'react'



function App() {
  const [count, setCount] = useState(0)

  return (
    
    <div className="App">
      <div className="card">
        <div className="card-body">
          <div className="">
            <h1 className="text-2xl font-bold">
              test1234
            </h1>
            <p className="text-muted">
              <no value>
            </p>
            <div className="button-group">
              <button
                onClick={() => setCount((count) => count + 1)}
                className="btn btn-primary"
              >
                Count is {count}
              </button>
              
            </div>
          </div>
        </div>
      </div>
    </div>
    
  )
}

export default App
