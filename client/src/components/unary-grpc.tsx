import { createSignal } from 'solid-js'
import { createConnectTransport, createPromiseClient } from '@bufbuild/connect-web'
import { GreetService } from '../proto/greet/greet_connectweb'

const transport = createConnectTransport({
  baseUrl: import.meta.env.API_GRPC
})

const client = createPromiseClient(GreetService, transport)

const [greeting, setGreeting] = createSignal('Greeting message!')

let name = ''

async function sendMessage() {
  const res = await client.greet({
    name
  })

  setGreeting(res.greeting)  
}

export default function() {
  return (
    <div>
      <div>
        <input 
          placeholder="Enter your name" 
          oninput={(event) => name = (event.target as HTMLInputElement).value } 
        />
        &nbsp;
        <button 
          onclick={() => sendMessage()}
        >
          Send
        </button>
      </div>
      <br />
      <div>
        {greeting()}
      </div>
    </div>
  )
}
