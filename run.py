import paramiko

host = "192.168.0.11"
port = 22
username = "root"
password = "m"

client = paramiko.SSHClient()
client.set_missing_host_key_policy(paramiko.AutoAddPolicy())
try:
    client.connect(hostname=host, port=port,  username=username, password=password)
    print("Successfully connected to SSH server.")
except paramiko.AuthenticationException:
    print("Authentication failed. Please check your username and password.")
except paramiko.SSHException as e:
    print(f"SSH connection error: {e}")
except Exception as e:
    print(f"An unexpected error occurred: {e}")

if client.get_transport().is_active(): # Check if connection is still active
    stdin, stdout, stderr = client.exec_command('zfs list')
    print("Command output:")
    print(stdout.read().decode())
    if stderr.read().decode():
        print("Error output:")
        print(stderr.read().decode())
