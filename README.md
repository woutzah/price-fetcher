# build image
> docker build --tag price-fetcher:1 .

# run built image
> docker run -p 3000:3000 -d price-fetcher:1