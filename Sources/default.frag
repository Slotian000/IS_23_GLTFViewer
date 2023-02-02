#version 410 core

out vec4 FragColor;

//in vec3 color;
in vec2 TexCoord;

uniform sampler2D ourTexture;

void main()
{
    FragColor = texture(ourTexture, TexCoord);
    //FragColor = FragColor * vec4(color,0.3f);
    //FragColor = vec4(color, 1.0f);
}