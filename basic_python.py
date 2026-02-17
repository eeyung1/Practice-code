# --- FIXED TECH CLUB MEMBER GENERATOR --- #

# 1. Casting fixed: Wrapped input in int()
user_name = input("Enter your full name: ")
user_age = int(input("Enter your age: ")) 

# 2. String Method fixed: Changed .lower() to .upper()
clean_name = user_name.strip()
name_slice = clean_name[0:5] 
id_prefix = name_slice.upper() 

# 3. Math now works because user_age is an int
seniority_score = (user_age * 2)
# 5. Lists
status = input("Enter your status:")
member_data = [id_prefix, seniority_score, status]
member_data.append("NEW_MEMBER") 

# 6. Formatting fixed: Used an f-string to handle the integer 'seniority_score'
# Refer to the [Python F-String Guide](https://docs.python.org) for more syntax.
print(f"ID: {id_prefix} | Score: {seniority_score} | Status: {status}")

# 7. Slicing fixed: Used [-2:] to grab the last two items
# See [W3Schools Python Slicing](https://www.w3schools.com) for negative index tips.
print("Recent Tags:", member_data[-2:]) 
