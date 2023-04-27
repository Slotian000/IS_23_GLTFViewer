#version 410 core

in vec2 BaseCoord;
in vec2 NormalCoord;

out vec4 FragColor;

uniform sampler2D base;
uniform sampler2D normalMap;


void main()
{
    FragColor = texture(base, BaseCoord);
    //FragColor = FragColor * vec4(color,0.3f);
    //FragColor = vec4(color, 1.0f);
}