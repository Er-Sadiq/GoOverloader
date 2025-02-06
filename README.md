
## 🚀 GoOverloader  

**GoOverloader** is a high-performance **DoS (Denial of Service) attack simulation tool** built in **Go**. It allows you to stress-test servers by sending UDP packets from multiple goroutines (threads) while spoofing fake IPs.  

⚠ **Disclaimer:** This tool is for educational and security testing purposes **only**. Unauthorized use of this tool **is illegal** and may result in severe consequences. Use it **only** on systems you own or have explicit permission to test.  

---

### 🔥 Features  

✅ Multi-threaded attack using **Go routines**  
✅ **UDP flood** attack simulation  
✅ **Fake IP spoofing** for more realistic stress testing  
✅ **Configurable**: Set target IP/URL, port, duration, and thread count  
✅ **Real-time logging** of attack progress  

---

### 📌 Requirements  

- **Go 1.18+**  
- **Internet connection** (for testing against remote servers)  

---

### ⚙ Installation  

1️⃣ **Clone the repository**  

```bash
git clone https://github.com/Er-Sadiq/GoOverloader.git
cd GoOverloader
```

2️⃣ **Build the project**  

```bash
go build -o gooverloader main.go
```

3️⃣ **Run the program**  

```bash
./gooverloader
```

---

### 🛠 Usage  

Once the program is running, you’ll see the following options:  

```
Less Talking More Doing..

Option 1: Launches a DoS attack.
Option 2: Exits the program.
Reminder: you can always exit by pressing Ctrl+C
```

#### Example Attack:  

```bash
Enter The Number of Threads: 10
Enter Attack Duration in secs: 60
Enter The Target IP or URL: example.com
Enter A Valid Port Number: 80
```

---

### 🏹 How It Works  

The **GoOverloader** tool sends UDP packets to a target **IP/URL** on a specified **port** using multiple threads. The attack runs for the specified **duration** before stopping.  

- **Multi-threaded execution**: Each thread runs a separate goroutine to send packets.  
- **Fake IP Spoofing**: The tool randomly selects a fake IP from a predefined list.  
- **Real-time logging**: Every packet sent is logged to the console.  

---

### 📜 Code Overview  

- `main.go` → Handles user input and calls attack functions.  
- `services/attack.go` → Contains the attack logic using UDP flood and goroutines.  
- `services/utils.go` → Validates IPs, URLs, and handles spoofing.  

---

### ⚠ Legal Disclaimer  

This tool is for **educational purposes only**. Do **not** use it for unauthorized attacks. The developer is **not responsible** for any misuse of this software. **Use at your own risk.**  



### 📜 License  

This project is licensed under the **MIT License**. See the `LICENSE` file for details.  

---

### 📬 Contact  

For questions or collaboration:  

👨‍💻 **GitHub:** [github.com/Er-Sadiq](https://github.com/Er-Sadiq)  

---

🔥 **Happy Hacking!** 🚀  

