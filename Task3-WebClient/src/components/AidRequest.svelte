<script lang="ts">
	import * as Drawer from '$lib/components/ui/drawer';
	import { Button } from '$lib/components/ui/button';
	import type { HelpRequest } from '../types/types';

	export let request: HelpRequest;
</script>

<Drawer.Content>
	<div class="mx-auto w-full max-w-xl">
		<Drawer.Header>
			<Drawer.Title>Детальний опис запиту</Drawer.Title>
			<Drawer.Description
				>Якщо ви бачите тут неумісний контент, негайно зверніться до адміністратора</Drawer.Description
			>
		</Drawer.Header>
		<div class="h-[240px] p-4 pb-0">
			<p class="mb-4 text-lg font-semibold">{request.title}</p>
			<p class="mb-4 text-gray-600">{request.description}</p>
		</div>
		<Drawer.Footer>
			<Drawer.Close asChild let:builder>
				<Button builders={[builder]} variant="outline">Назад</Button>
			</Drawer.Close>
		</Drawer.Footer>
	</div>
</Drawer.Content>
