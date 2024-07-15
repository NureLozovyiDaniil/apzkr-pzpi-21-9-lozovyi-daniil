<script lang="ts">
	import { onMount } from 'svelte';
	import { Button } from '$lib/components/ui/button';
	import { toast } from 'svelte-sonner';
	import { Input } from '$lib/components/ui/input';
	import { Label } from '$lib/components/ui/label';
	import { Textarea } from '$lib/components/ui/textarea';
	import * as Drawer from '$lib/components/ui/drawer';
	import {
		Table,
		TableBody,
		TableCell,
		TableHead,
		TableHeader,
		TableRow
	} from '$lib/components/ui/table';
	import { Checkbox } from '$lib/components/ui/checkbox';
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu'; // Assuming you have a Dropdown component
	import { MoreHorizontal, Filter, Plus, Trash2, Text } from 'lucide-svelte';
	import type { HelpRequest } from '../../types/types';
	import { Statuses, getCookie } from '../../types/types';
	import AidRequest from '../../components/AidRequest.svelte';

	let requests: HelpRequest[] = [];
	let filterValue = '';
	let reqDetailId = '';
	let openDetails = false;
	let selectedStatuses: Set<string> = new Set();

	onMount(getRequests);

	async function getRequests() {
		try {
			const response = await fetch('http://localhost:5001/api/v1/help_requests', {
				headers: {
					Authorization: `Bearer ${getCookie('authToken')}`
				}
			});

			let data = await response.json();
			console.log(data);

			requests = data;
		} catch (error) {
			console.error('Error fetching tasks:', error);
		}
	}

	function reqById(id: string) {
		return requests.find((req) => req.id === id)!;
	}

	let title = '';
	let description = '';

	async function createRequest() {
		console.log('Creating request with title:', title, 'and description:', description);
		const response = await fetch('http://localhost:5001/api/v1/help_requests', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
				Authorization: `Bearer ${getCookie('authToken')}`
			},
			body: JSON.stringify({ title, description })
		});

		if (response.ok) {
			getRequests();
			toast.success('Request created successfully');
		} else {
			console.error('Failed to create request');
			toast.error('Failed to create request', {
				description: response.status + ' ' + (await response.text())
			});
		}
	}

	async function deleteRequest(reqId: string) {
		console.log('Deleting request with id:', reqId);
		const response = await fetch(`http://localhost:5001/api/v1/help_requests/${reqId}`, {
			method: 'DELETE',
			headers: {
				Authorization: `Bearer ${getCookie('authToken')}`
			}
		});

		if (response.ok) {
			getRequests();
			toast.success('Request deleted successfully');
		} else {
			console.error('Failed to delete request');
			toast.error('Failed to delete request', {
				description: response.status + ' ' + (await response.text())
			});
		}
	}

	$: filteredRequests = requests.filter(
		(req) =>
			req.title.toLowerCase().includes(filterValue.toLowerCase()) &&
			(selectedStatuses.size === 0 || selectedStatuses.has(req.status))
	);

	function handleReqClick(reqId: string) {
		reqDetailId = reqId;
		toggleDetails();
		console.log(`Navigating to task ${reqId}`);
	}

	function toggleDetails() {
		openDetails = !openDetails;
	}

	function toggleStatus(status: string) {
		if (selectedStatuses.has(status)) {
			selectedStatuses.delete(status);
		} else {
			selectedStatuses.add(status);
		}
		console.log(selectedStatuses);
	}
</script>

<div class="container mx-auto p-4">
	<h1 class="mb-2 text-3xl font-bold">Поточні запити</h1>
	<p class="mb-6 text-gray-600">Тут відображається список запитів на допомогу на поточний час.</p>

	<div class="mb-4 flex justify-between">
		<div class="flex w-1/3">
			<Input placeholder="Пошук за заголовком..." class="w-2/3" bind:value={filterValue} />
			<DropdownMenu.Root>
				<DropdownMenu.Trigger class="w-[180px]">
					<Button variant="ghost">Фільтрувати &nbsp; <Filter /></Button>
				</DropdownMenu.Trigger>
				<DropdownMenu.Content>
					<DropdownMenu.Group>
						{#each Statuses as status}
							<DropdownMenu.Item on:click={() => toggleStatus(status)}>
								<Checkbox checked={selectedStatuses.has(status)} />
								<span>&nbsp;{status}</span>
							</DropdownMenu.Item>
						{/each}
					</DropdownMenu.Group>
				</DropdownMenu.Content>
			</DropdownMenu.Root>
		</div>

		<div class="flex">
			<Drawer.Root>
				<Drawer.Trigger asChild let:builder>
					<Button builders={[builder]}><Plus />&nbsp; Створити запит</Button>
				</Drawer.Trigger>
				<Drawer.Content>
					<div class="mx-auto w-full max-w-xl">
						<Drawer.Header>
							<Drawer.Title>Створіть власний запит</Drawer.Title>
							<Drawer.Description>Детально вкажіть свої потреби у описі</Drawer.Description>
						</Drawer.Header>
						<div class="h-[240px] p-4 pb-0">
							<Label for="title">Заголовок</Label>
							<Input placeholder="Title" id="title" class="mb-3" bind:value={title} />
							<Label for="message">Опис</Label>
							<Textarea
								placeholder="Type your message here."
								id="message"
								bind:value={description}
							/>
						</div>
						<Drawer.Footer>
							<Button on:click={createRequest}>Створити</Button>
							<Drawer.Close asChild let:builder>
								<Button builders={[builder]} variant="outline">Назад</Button>
							</Drawer.Close>
						</Drawer.Footer>
					</div>
				</Drawer.Content>
			</Drawer.Root>
		</div>
	</div>

	<Table>
		<TableHeader>
			<TableRow>
				<TableHead class="w-[50px]"></TableHead>
				<TableHead>Заголовок</TableHead>
				<TableHead>Короткий опис</TableHead>
				<TableHead>Поточний статус</TableHead>
				<TableHead class="w-[50px]"></TableHead>
			</TableRow>
		</TableHeader>
		<TableBody>
			<Drawer.Root bind:open={openDetails}>
				<AidRequest request={reqById(reqDetailId)} />
			</Drawer.Root>
			{#each filteredRequests as req}
				<TableRow>
					<TableCell>
						<Checkbox />
					</TableCell>
					<TableCell>
						<button class="text-left hover:underline" on:click={() => handleReqClick(req.id)}>
							{req.title}
						</button>
						<!-- <p class="text-sm text-gray-500">{req.description}</p> -->
					</TableCell>
					<TableCell>{req.description}</TableCell>
					<TableCell>{req.status}</TableCell>
					<TableCell>
						<DropdownMenu.Root>
							<DropdownMenu.Trigger>
								<Button variant="ghost" size="icon">
									<MoreHorizontal class="h-4 w-4" />
								</Button>
							</DropdownMenu.Trigger>
							<DropdownMenu.Content>
								<DropdownMenu.Group>
									<DropdownMenu.Item
										on:click={() => handleReqClick(req.id)}
										class="flex items-center"
									>
										<Text />&nbsp;Детально
									</DropdownMenu.Item>
									{#if getCookie('userRole') === '4'}
										<DropdownMenu.Separator />
										<DropdownMenu.Item on:click={() => deleteRequest(req.id)} class="text-red-500"
											><Trash2 />&nbsp;Delete</DropdownMenu.Item
										>
									{/if}
								</DropdownMenu.Group>
							</DropdownMenu.Content>
						</DropdownMenu.Root>
					</TableCell>
				</TableRow>
			{/each}
		</TableBody>
	</Table>
</div>
