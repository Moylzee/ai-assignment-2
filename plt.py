import pandas as pd
import matplotlib.pyplot as plt

data = pd.read_csv('fitness_data.csv')

fig, axs = plt.subplots(3, 1, figsize=(10, 12))

axs[0].plot(data['Generation'], data['AverageFitness'], label='Average Fitness', color='blue')
axs[0].set_title('Fitness Progression Over Generations')
axs[0].set_xlabel('Generation')
axs[0].set_ylabel('Average Fitness')
axs[0].grid(True)
axs[0].legend()

axs[1].bar(data['Generation'], data['TotalCooperations'], width=0.4, label='Cooperations', align='center', color='green')
axs[1].bar(data['Generation'], data['TotalDefects'], width=0.4, label='Defects', align='edge', color='red')
axs[1].set_title('Cooperations and Defects Per Generation')
axs[1].set_xlabel('Generation')
axs[1].set_ylabel('Count')
axs[1].legend()

data['Strategy'] = data['TotalDefects'] > 2500
axs[2].plot(data['Generation'], data['Strategy'], label='Cooperate (0) / Defect (1)', color='purple')
axs[2].set_title('Player Strategy (Cooperate/Defect) Over Generations')
axs[2].set_xlabel('Generation')
axs[2].set_ylabel('Strategy (0 = Cooperate, 1 = Defect)')
axs[2].grid(True)
axs[2].legend()

plt.tight_layout()
plt.show()
