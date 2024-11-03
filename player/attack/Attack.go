components {
  id: "attack"
  component: "/player/attack/attack.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"air_attack\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/pngs.atlas\"\n"
  "}\n"
  ""
  position {
    y: -3.0
  }
  scale {
    x: 1.3734
    y: 0.92911
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_TRIGGER\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"attack\"\n"
  "mask: \"stone\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_SPHERE\n"
  "    position {\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 1\n"
  "  }\n"
  "  data: 7.004796\n"
  "}\n"
  ""
}
