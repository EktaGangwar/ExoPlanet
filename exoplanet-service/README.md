# Exoplanet Microservice

This microservice allows users to manage a catalog of exoplanets. It supports adding, listing, updating, and deleting exoplanets, as well as estimating fuel requirements for travel.

## API Endpoints

- `POST /exoplanets` - Add a new exoplanet
- `GET /exoplanets` - List all exoplanets
- `GET /exoplanets/:id` - Get details of a specific exoplanet
- `PUT /exoplanets/:id` - Update an exoplanet
- `DELETE /exoplanets/:id` - Delete an exoplanet
- `GET /exoplanets/:id/fuel?crewCapacity={}` - Estimate fuel requirements

## Running the Service

1. Build and run the service:

```sh
docker build -t exoplanet-service .
docker run -p 8080:8080 exoplanet-service
