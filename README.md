# Swim Pack Tool
Desktop GUI tool for porting, modifying, and creating overlays for Minecraft Bedrock Edition.
<br><br>
<img width="485" alt="Screenshot 2025-02-24 100829" src="https://github.com/user-attachments/assets/86389037-b78d-47a3-a29b-460826d3aa19" />
# Porting
Converts 1.7.10 - 1.8.9 Java edition packs into Bedrock edition MCPACKs.
<br>
Upload Java edition pack locally or appropriate Mediafire download link, and then hit the Port button when ready.
<br>
A box can optionally be checked to disable the offset applied to make cubemaps appear better alligned.
<br>
You can specify in the cubemap override text input for which skybox the porter should attempt to find and convert into the ported bedrock's pack cubemap.
<br>
Eg. starfield01, sky01, night01, etc.

# Recoloring
Upload texture pack locally or appropriate Mediafire download link, and then hit the Recolor button when ready.
<br>
The choose color button takes you to a menu where you can pick any color, hex code colors supported as well.
<br><br>
<img width="485" alt="recolor" src="https://github.com/user-attachments/assets/e50715be-d486-45af-8fd9-e2f27d25a5d8" />
<br>
Three algorthims are supported via drop down menu selection. 
<br>
Tint applies a direct pixel color overlay on existing pixels.
<br>
Gray tint works the same but is intended for darker colors.
<br>
Hue works the best for any bright or vibrant colors while maintaining color balance.
![recolor](https://github.com/user-attachments/assets/bb00eacd-3bc9-4ffa-bd9a-5d6b6710c6ca)

# Rescaling
Upload texture pack locally or appropriate Mediafire download link, and then hit the Rescale button when ready.
<br>
Select what scale you want to change the pack to via drop down
<br>
Select which algorithm you want to use, there are a lot of possible ones, so play around.
<br>
This feature is most useful for scaling down 128x+ texture packs to 64x and 32x to save storage, performance, and in some cases get a more pixel art styled design.

# Sky overlay
Converts any image into a cubemap overlay for bedrock edition. This uses equirectangular projection to appear as seamless as possible in game.
<br>
Select the image you want, and then optionally with the slider modify the blend value, which is how much pixel blending to apply along touching edges of each face of the cubemap.
<br>
You can optionally disable the skybox offset if you like having your skies look unalligned for some reason.
<br>
Higher aspect ratio images yield far better results. How rectungular or how square shaped the image is will impact blending and quality.

![sky](https://github.com/user-attachments/assets/d3eeb55e-9c4f-4aa5-8fd8-701042d3708a)


# Java sky porter
Converts any bedrock cubemap into a single image for a java edition sky. This is mostly useful for taking the outputs of the Sky overlay and putting it into java edition format.

# Inventory maker
Converts any image or gif into an inventory overlay texture pack. Gif sizes will apply and after a set amount of frames it will be cut off due to bedrock spritesheet limitations.
<br><br>
![184106224-eddd3236-c07c-4de4-9d8b-020231f2077d](https://github.com/user-attachments/assets/26bfc737-141f-4937-b710-625413065909)

# Crosshair maker
Converts any image into a crosshair overlay texutre pack. The image will be scaled down and transparent in appearance due to how mcpe renders crosshairs.

# Misc fixers
This tool has some other small features for attempting to fix font, skies, and particles in existing packs. These can be very hit or miss. 
<br>
This simply won't work if the pack lacks somethin crucial like font or particles.
