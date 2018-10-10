# README

## System Requirements
( Golang:1.11, Docker )
# Run
### `$ docker-compose up`
##### End points:
- `http://localhost:8080/` Welcome
- `http://localhost:8080/register` Register a user
we have 2 types of registeration, 
1- The patient type with field (userType=0):
` http -v POST http://localhost:8080/register firstName=fpatient lastName=lpatient email=fake@fake password=123456 userType:=0 `
2- The doctor type with field (userType=1):
`http -v POST http://localhost:8080/register firstName=fdoctor lastName=ldoctor email=faked@faked password=123456 userType:=1`
- `http://localhost:8080/reserve` Reserve an appoinment
`http -v POST http://localhost:8080/reserve doctor:=1 patient:=95 notes=whatever timePoint=2018-11-27T00:00:00Z`
- `http://localhost:8080/patient/:id` Find a patient with id
- `http://localhost:8080/patients` List all the patients
- `http://localhost:8080/doctor/:id` Find a doctor with id
- `http://localhost:8080/doctors` List all the doctors
- `http://localhost:8080/appoinment/doctor/:id` List all the appoinments for a doctor with id
# TODO
- #### Adding mocks for testing

