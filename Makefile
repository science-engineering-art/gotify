
vendor:
	cd api && go mod tidy && go mod vendor && cd ..
	cd peer && go mod tidy && go mod vendor && cd .. 
	cd tracker && go mod tidy && go mod vendor && cd .. 