# go-statx

## Installation


### The the CLI

`go get -u github.com/mdp/go-statx/...`

### Just the library

`go get -u github.com/mdp/go-statx`

## Usage


### CLI

`statx login +14158675309`  
Enter verification code from StatX and if successful you will receive API keys

Get a list of groups  
`statx list --apikey 12345abcdef --authtoken yourauthtoken`

Get a list of Stats in a group  
`statx list --apikey 12345abcdef --authtoken yourauthtoken --group groupid`

Update the value of a Stat  
`statx update --apikey 12345abcdef --authtoken yourauthtoken --group groupid --stat statid --value 47`

### Library

```go
client := statx.NewAuthenticatedClient(nil, "apikey", "authtoken")
statList, _, err := client.Stats.List("GroupID")

# Update a stat
stat := &statx.Stat{Value: "47"}
updatedStat, _, err := client.Stats.Update("GroupID", statList[0].ID, stat)
```

