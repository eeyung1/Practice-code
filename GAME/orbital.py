import pygame
import math

# settings
WIDTH, HEIGHT = 800, 600
FPS = 60
G = 0.5  # gravity constant

# sun
sun = {
    "x": WIDTH // 2,
    "y": HEIGHT // 2,
    "mass": 1000,
    "radius": 20,
    "color": (255, 255, 0)  # yellow
}

# planet
planet = {
    "x": WIDTH // 2 + 200,
    "y": HEIGHT // 2,
    "mass": 1,
    "radius": 10,
    "color": (0, 100, 255),  # blue
    "vel_x": 0,  # needs to move UP to start orbiting
    "vel_y": -1.58   # what value makes it orbit?
}

pygame.init()
screen = pygame.display.set_mode((WIDTH, HEIGHT))
pygame.display.set_caption("Orbital Simulation")
clock = pygame.time.Clock()

running = True
while running:
    for event in pygame.event.get():
        if event.type == pygame.QUIT:
            running = False

    # calculate gravity pull towards sun
    dx = sun["x"] - planet["x"]
    dy = sun["y"] - planet["y"]
    distance = math.sqrt(dx**2 + dy**2)
    force = G * sun["mass"] / distance**2

    # apply force to planet velocity
    planet["vel_x"] += force * dx / distance
    planet["vel_y"] += force * dy / distance

    # move planet
    planet["x"] += planet["vel_x"]
    planet["y"] += planet["vel_y"]

    # draw
    screen.fill((0, 0, 10))  # dark space background
    pygame.draw.circle(screen, sun["color"], (int(sun["x"]), int(sun["y"])), sun["radius"])
    pygame.draw.circle(screen, planet["color"], (int(planet["x"]), int(planet["y"])), planet["radius"])
    pygame.display.flip()

    clock.tick(FPS)

pygame.quit()