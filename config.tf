provider "example" {
  api_key = "s3cur3t0k3n=="
  url     = "https://api.example.org"
}

resource "example_server" "my-first-server" {
  cpus     = 2
  hostname = "superserver9000"
}
