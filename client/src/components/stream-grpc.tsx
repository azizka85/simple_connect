import { createConnectTransport, createPromiseClient } from '@bufbuild/connect-web'
import { TimingOutputService } from '../proto/greet/timing_connectweb'

const transport = createConnectTransport({
  baseUrl: import.meta.env.API_GRPC
})

const client = createPromiseClient(TimingOutputService, transport)

let message = ''

async function accept() {
  const res = client.accept({
    message
  })

  for await (const data of res) {
    console.log(
      'client.send: ',
      data
    )
  }
}

export default function() {
  return (
    <div>
      <div>
        <input 
          placeholder="Enter message" 
          oninput={(event) => message = (event.target as HTMLInputElement).value }
        />
        &nbsp;
        <button 
          onclick={() => accept()}
        >
          Accept
        </button>
      </div>
      <br />
      <div>
        
      </div>
    </div>
  )
}
