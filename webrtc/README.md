#Dependencies
1. [Node.js](http://nodejs.org): Use the installer provided on the NodeJS

#Setup
Change to the easyrtc folder an then install node modules locally
`sudo npm install`

#Running
You can run easyrtc either from console or as a service.

- From console
   1. Change to the easyrtc application folder.
   2. Run the server using the node command: 
      - node server.js

- As a service
   1. Create a file named easyrtc.conf an save it in /etc/init.
   2. Open the document and copy the following script:

	# Saves log to /var/log/upstart/easyrtc.log
	console log

	# Starts only after drives are mounted.
	start on started mountall

	stop on shutdown

	# Automatically Respawn. But fail permanently if it respawns 10 times in 5 seconds:
	respawn
	respawn limit 10 5

	script
	    # exec [path to the binary of node] [path where easyrtc server is located]
	    exec /opt/nod-v0.10.28/node /home/aaron/go/kails/src/bitbucket.com/abijr/kails/webrtc/easyrtc/server.js
	end script 

   3. Open the console and type: 
      - Start server
        - sudo start easyrtc
      - Stop server
        - sudo stop easyrtc





	


	
