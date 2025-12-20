<script>
	import { Canvas } from '@threlte/core';
	import { T } from '@threlte/core';
	import { OrbitControls, Environment, ContactShadows } from '@threlte/extras';
	import ControllerModel from './controller-model.svelte';

	// Importăm componenta creată mai sus
</script>

<div class="scene-wrapper">
	<Canvas>
		<T.PerspectiveCamera makeDefault position={[0, 5, 10]} fov={25}>
			<OrbitControls enableZoom={true} enablePan={false} autoRotate={true} />
		</T.PerspectiveCamera>

		<T.AmbientLight intensity={0.5} />
		<T.DirectionalLight position={[5, 10, 5]} intensity={2} castShadow />

		<ControllerModel />

		<ContactShadows opacity={0.5} blur={2} scale={10} far={2} />
	</Canvas>
</div>

<style>
	.scene-wrapper {
		width: 100%;
		height: 300px;
		background: transparent;
	}
</style>
