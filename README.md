### Dependencies installation:

1. Install node & gulp globally
2. Install all node dependencies: `cd client && npm i`
3. Install golang 1.9.2 or greater
4. In .zshrc or .bashrc or your profile add this 2 lines:
    
       export GOPATH=$HOME/go
       export PATH=$PATH:$GOPATH/bin
5. close and open new terminal
       
    
### Developing on Mac

* Reload go code:
  
  1.  Style codegansta gin for autoreload of golang: `go get github.com/codegangsta/gin`
  2.  `sh go-reload.sh` 
      
      Output:
          
          [gin] Listening on port 3000
          [gin] Building...
          [gin] Build finished
  3. Go to localhost:3000         

* Reload client code:
     
  1. Open new terminal at project root
  2. `sh client-reload.sh` 

