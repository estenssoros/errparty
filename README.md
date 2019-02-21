# errparty

because errgroup never seems to work like it should

## Getting Started

```
import (
  "time"
  "log"
  "github.com/estenssoros/errparty"
)

func main(){
  party := errparty.ErrParty{}
  party.RSVP(func() error{
    time.Sleep(5 * time.Second)
    return nil
  })

  party.RSVP(func() error{
    if err := myFunc(); err != nil{
      return err
    }
    return nil
  })

  if err := party.Party(); err != nil{
    log.Fatal(err)
  }
}
```
