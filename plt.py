import pandas as pd
import matplotlib.pyplot as plt

# Load the fitness data from the CSV file
data = pd.read_csv('fitness_data.csv')

# Plot the fitness data
plt.figure(figsize=(10, 6))
plt.plot(data['Generation'], data['AverageFitness'], label='Average Fitness', color='blue')

# Add titles and labels
plt.title('Fitness Progression Over Generations')
plt.xlabel('Generation')
plt.ylabel('Average Fitness')
plt.grid(True)
plt.legend()

# Show the plot
plt.show()
