import pandas as pd
import matplotlib.pyplot as plt

# Load the fitness data from the CSV file
data = pd.read_csv('fitness_data.csv')

# Create subplots with 3 rows and 1 column
fig, axs = plt.subplots(3, 1, figsize=(10, 12))

# 1st Subplot: Average Fitness Progression
axs[0].plot(data['Generation'], data['AverageFitness'], label='Average Fitness', color='blue')
axs[0].set_title('Fitness Progression Over Generations')
axs[0].set_xlabel('Generation')
axs[0].set_ylabel('Average Fitness')
axs[0].grid(True)
axs[0].legend()

# 2nd Subplot: Cooperations and Defects Bar Chart
axs[1].bar(data['Generation'], data['TotalCooperations'], width=0.4, label='Cooperations', align='center', color='green')
axs[1].bar(data['Generation'], data['TotalDefects'], width=0.4, label='Defects', align='edge', color='red')
axs[1].set_title('Cooperations and Defects Per Generation')
axs[1].set_xlabel('Generation')
axs[1].set_ylabel('Count')
axs[1].legend()

# 3rd Subplot: Strategy (Cooperate/Defect) over Generations
# Let's create a binary "strategy" column where 1 = Defect and 0 = Cooperate
# Assuming a generation where the player defects more than 2500 times is considered as defecting, otherwise cooperating
data['Strategy'] = data['TotalDefects'] > 2500  # True if defecting, False if cooperating
axs[2].plot(data['Generation'], data['Strategy'], label='Cooperate (0) / Defect (1)', color='purple')
axs[2].set_title('Player Strategy (Cooperate/Defect) Over Generations')
axs[2].set_xlabel('Generation')
axs[2].set_ylabel('Strategy (0 = Cooperate, 1 = Defect)')
axs[2].grid(True)
axs[2].legend()

# Adjust layout
plt.tight_layout()

# Show the plot
plt.show()
