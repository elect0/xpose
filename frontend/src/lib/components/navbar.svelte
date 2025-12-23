<script lang="ts">
	import { page } from '$app/state';

	$effect(() => {
		console.log(page.url.pathname);
	});
</script>

<nav
	class="fixed top-0 right-0 left-0 z-50 w-full bg-transparent py-1 md:py-3"
>
	<div class="container mx-auto px-4 sm:px-6 lg:px-8">
		<div class="flex h-16 items-center justify-between">
			<div class="flex-shrink-0">
				<a href="/" class="itmes-center flex space-x-3">
					<div class="text-[#117CD9]] text-2xl font-bold md:text-3xl">LI.</div>
				</a>
			</div>
			<!-- <div class="space-x-3"> -->
			<!-- 	<a -->
			<!-- 		href="/the-loop" -->
			<!-- 		class={page.url.pathname.startsWith('/the-loop') ? 'text-[#F25C69]' : 'text-[#0C0202]'} -->
			<!-- 		>The Loop</a -->
			<!-- 	> -->
			<!-- 	<a -->
			<!-- 		href="/packs" -->
			<!-- 		class={page.url.pathname.startsWith('/packs') ? 'text-[#F25C69]' : 'text-[#0C0202]'} -->
			<!-- 		>Packs</a -->
			<!-- 	> -->
			<!-- </div> -->
			<button
				class="transform cursor-pointer rounded-full border-2 border-black bg-[#FFFBEF] px-4 py-2 text-sm font-semibold text-black shadow-[2px_2px_0_0_#000] transition-all hover:translate-x-1 hover:translate-y-1 hover:shadow-none active:translate-x-1 active:translate-y-1 active:shadow-none md:text-base"
			>
				Ready up!
			</button>
		</div>
	</div>
</nav>
