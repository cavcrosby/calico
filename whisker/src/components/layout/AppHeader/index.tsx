import { CalicoCatIcon, CalicoWhiskerIcon } from '@/icons';
import { Box, Flex, Heading, Link, Text } from '@chakra-ui/react';
import React from 'react';
import { appHeaderStyles } from './styles';

const AppHeader: React.FC = () => {
    return (
        <Box>
            <Flex as='header' sx={appHeaderStyles}>
                <Flex alignItems='center'>
                    <CalicoWhiskerIcon fontSize='xl' />
                    <Heading fontSize='2xl'>Calico Whisker</Heading>
                </Flex>

                <Flex alignItems='center' gap={2}>
                    <Flex alignItems='flex-end' flexDirection='column' gap={0}>
                        <Text fontSize='xs' as='span' fontWeight='medium'>
                            Calico Whisker is a simplified version of the
                        </Text>
                        <Link
                            fontSize='xs'
                            fontWeight='bold'
                            isExternal
                            href='https://calicocloud.io'
                        >
                            Service Graph from Calico Cloud
                        </Link>
                    </Flex>

                    <CalicoCatIcon fontSize='35px' />
                </Flex>
            </Flex>
        </Box>
    );
};

export default AppHeader;
