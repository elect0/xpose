<script>
  import { T, useTask } from '@threlte/core';
  import { GLTF } from '@threlte/extras';

  let rotation = 0;

  // Aici useTask funcționează perfect
  useTask((delta) => {
    rotation += delta * 0.5;
  });
</script>

<T.Group rotation.y={rotation}>
  <GLTF 
    url="/models/computer.glb" 
    position={[0, 0, 0]} 
    scale={1} 
  />
</T.Group>
