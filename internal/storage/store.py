import sqlite3

def store(data):
    # Connect to the database
    conn = sqlite3.connect("data.db")
    c = conn.cursor()

    # Create the table if it doesn't exist
    c.execute('''CREATE TABLE IF NOT EXISTS data (timestamp text, heart_rate int, steps int)''')

    # Insert the data into the table
    c.execute("INSERT INTO data VALUES (?, ?, ?)", (data.Timestamp, data.HeartRate, data.Steps))

    # Commit the transaction
    conn.commit()

    # Close the connection
    conn.close()
