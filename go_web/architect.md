architect building api with golang

# Flow 
- package pkg store a file to connect with database 
- internal as a core to process business logic must have at least 4 layers:
   - entity is the lowest layer 
   - repository layer to store operations with the database using the entity layer
   - usecase layer as the service handle business logic 
   - handler layer - highest layer (controller) to create an api for the app
