import {RiShareLine} from 'react-icons/ri';
import {defineField} from 'sanity';

const socialFields = defineField({
	title: 'Social',
	name: 'socialFields',
	type: 'object',
	icon: RiShareLine,
	fields: [
		{
			name: 'twitter',
			type: 'url',
			title: 'Twitter URL'
		},
		{
			name: 'github',
			type: 'url',
			title: 'Github URL'
		},
		{
			name: 'discord',
			type: 'url',
			title: 'Discord URL'
		},
		{
			name: 'linkedin',
			type: 'url',
			title: 'LinkedIn URL'
		}
		,
		{
			name: 'crunchbase',
			type: 'url',
			title: 'Crunchbase URL'
		}
	]
});

export default socialFields;
