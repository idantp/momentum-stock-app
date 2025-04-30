package db

// ConnectDB establishes a connection to the PostgreSQL database using the provided connection string.
// func ConnectDB() (*pgx.Conn, error) {
// 	// Load environment variables from .env file
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}

// 	// Get the PostgreSQL connection string from environment variables
// 	connStr := os.Getenv("DATABASE_URL")
// 	if connStr == "" {
// 		return nil, fmt.Errorf("DATABASE_URL not set in .env file")
// 	}

// 	// Connect to the PostgreSQL database
// 	conn, err := pgx.Connect(context.Background(), connStr)
// 	if err != nil {
// 		return nil, fmt.Errorf("unable to connect to database: %v", err)
// 	}

// 	return conn, nil
// }
