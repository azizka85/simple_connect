import StreamGRPC from './components/stream-grpc'
import UnaryGRPC from './components/unary-grpc'

export default function() {
  return (
    <>
      <h3>Unary gRPC type</h3>
      <UnaryGRPC />
      <hr />
      <h3>Stream gRPC type</h3>
      <StreamGRPC />
      <StreamGRPC />
    </>
  )
}
