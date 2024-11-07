components {
  id: "attack"
  component: "/player/attacks/attack.script"
  properties {
    id: "speed"
    value: "200.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "lifetime"
    value: "3.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "hit_value"
    value: "100.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "aoe"
    value: "true"
    type: PROPERTY_TYPE_BOOLEAN
  }
}
components {
  id: "ability_animation"
  component: "/player/attacks/ability_animation.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"ability\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/pngs.atlas\"\n"
  "}\n"
  ""
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
  "  data: 8.1866665\n"
  "}\n"
  ""
}
