# pip3 install pygame --break-system-packages

import pygame

# settings
WIDTH, HEIGHT = 700, 600
FPS = 60
GRAVITY = 0.1
BOUNCE = 0.8

# ball properties
ball_x = 400
ball_y = 100
ball_radius = 20
vel_x = 4
vel_y = 0

# pygame setup
pygame.init()
screen = pygame.display.set_mode((WIDTH, HEIGHT))
pygame.display.set_caption("Bouncing Ball")
clock = pygame.time.Clock()

running = True
while running:
    # 1. handle input
    for event in pygame.event.get():
        if event.type == pygame.QUIT:
            running = False

        # click to throw ball
        if event.type == pygame.MOUSEBUTTONDOWN:
            mouse_x, mouse_y = pygame.mouse.get_pos()
            ball_x = mouse_x
            ball_y = mouse_y
            vel_x = 4
            vel_y = -10  # negative = upward

    # 2. update physics
    vel_y += GRAVITY
    ball_x += vel_x
    ball_y += vel_y

    # 3. collision detection
    if ball_x + ball_radius >= WIDTH or ball_x - ball_radius <= 0:
        vel_x = -vel_x

    if ball_y + ball_radius >= HEIGHT:
        ball_y = HEIGHT - ball_radius
        vel_y = -vel_y * BOUNCE

    # 4. draw everything
    screen.fill((0, 0, 0))
    pygame.draw.circle(screen, (255, 0, 0), (int(ball_x), int(ball_y)), ball_radius)
    pygame.display.flip()

    clock.tick(FPS)

pygame.quit()
