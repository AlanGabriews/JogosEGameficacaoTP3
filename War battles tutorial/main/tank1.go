components {
  id: "tank1"
  component: "/main/tank1.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"tank-down\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/sprites.atlas\"\n"
  "}\n"
  ""
}
