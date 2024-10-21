# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from mgmt.v1alpha1 import connection_pb2 as mgmt_dot_v1alpha1_dot_connection__pb2


class ConnectionServiceStub(object):
    """Service for managing datasource connections.
    This is a primary data model in Neosync and is used in reference when hooking up Jobs to synchronize and generate data.
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetConnections = channel.unary_unary(
                '/mgmt.v1alpha1.ConnectionService/GetConnections',
                request_serializer=mgmt_dot_v1alpha1_dot_connection__pb2.GetConnectionsRequest.SerializeToString,
                response_deserializer=mgmt_dot_v1alpha1_dot_connection__pb2.GetConnectionsResponse.FromString,
                _registered_method=True)
        self.GetConnection = channel.unary_unary(
                '/mgmt.v1alpha1.ConnectionService/GetConnection',
                request_serializer=mgmt_dot_v1alpha1_dot_connection__pb2.GetConnectionRequest.SerializeToString,
                response_deserializer=mgmt_dot_v1alpha1_dot_connection__pb2.GetConnectionResponse.FromString,
                _registered_method=True)
        self.CreateConnection = channel.unary_unary(
                '/mgmt.v1alpha1.ConnectionService/CreateConnection',
                request_serializer=mgmt_dot_v1alpha1_dot_connection__pb2.CreateConnectionRequest.SerializeToString,
                response_deserializer=mgmt_dot_v1alpha1_dot_connection__pb2.CreateConnectionResponse.FromString,
                _registered_method=True)
        self.UpdateConnection = channel.unary_unary(
                '/mgmt.v1alpha1.ConnectionService/UpdateConnection',
                request_serializer=mgmt_dot_v1alpha1_dot_connection__pb2.UpdateConnectionRequest.SerializeToString,
                response_deserializer=mgmt_dot_v1alpha1_dot_connection__pb2.UpdateConnectionResponse.FromString,
                _registered_method=True)
        self.DeleteConnection = channel.unary_unary(
                '/mgmt.v1alpha1.ConnectionService/DeleteConnection',
                request_serializer=mgmt_dot_v1alpha1_dot_connection__pb2.DeleteConnectionRequest.SerializeToString,
                response_deserializer=mgmt_dot_v1alpha1_dot_connection__pb2.DeleteConnectionResponse.FromString,
                _registered_method=True)
        self.IsConnectionNameAvailable = channel.unary_unary(
                '/mgmt.v1alpha1.ConnectionService/IsConnectionNameAvailable',
                request_serializer=mgmt_dot_v1alpha1_dot_connection__pb2.IsConnectionNameAvailableRequest.SerializeToString,
                response_deserializer=mgmt_dot_v1alpha1_dot_connection__pb2.IsConnectionNameAvailableResponse.FromString,
                _registered_method=True)
        self.CheckConnectionConfig = channel.unary_unary(
                '/mgmt.v1alpha1.ConnectionService/CheckConnectionConfig',
                request_serializer=mgmt_dot_v1alpha1_dot_connection__pb2.CheckConnectionConfigRequest.SerializeToString,
                response_deserializer=mgmt_dot_v1alpha1_dot_connection__pb2.CheckConnectionConfigResponse.FromString,
                _registered_method=True)
        self.CheckConnectionConfigById = channel.unary_unary(
                '/mgmt.v1alpha1.ConnectionService/CheckConnectionConfigById',
                request_serializer=mgmt_dot_v1alpha1_dot_connection__pb2.CheckConnectionConfigByIdRequest.SerializeToString,
                response_deserializer=mgmt_dot_v1alpha1_dot_connection__pb2.CheckConnectionConfigByIdResponse.FromString,
                _registered_method=True)
        self.CheckSqlQuery = channel.unary_unary(
                '/mgmt.v1alpha1.ConnectionService/CheckSqlQuery',
                request_serializer=mgmt_dot_v1alpha1_dot_connection__pb2.CheckSqlQueryRequest.SerializeToString,
                response_deserializer=mgmt_dot_v1alpha1_dot_connection__pb2.CheckSqlQueryResponse.FromString,
                _registered_method=True)


class ConnectionServiceServicer(object):
    """Service for managing datasource connections.
    This is a primary data model in Neosync and is used in reference when hooking up Jobs to synchronize and generate data.
    """

    def GetConnections(self, request, context):
        """Returns a list of connections associated with the account
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetConnection(self, request, context):
        """Returns a single connection
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CreateConnection(self, request, context):
        """Creates a new connection
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UpdateConnection(self, request, context):
        """Updates an existing connection
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DeleteConnection(self, request, context):
        """Removes a connection from the system.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def IsConnectionNameAvailable(self, request, context):
        """Connections have friendly names, this method checks if the requested name is available in the system based on the account
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CheckConnectionConfig(self, request, context):
        """Checks if the connection config is connectable by the backend.
        Used mostly to verify that a connection is valid prior to creating a Connection object.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CheckConnectionConfigById(self, request, context):
        """Checks if the connection id is connectable by the backend.
        Used to verify that a connection is still connectable.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CheckSqlQuery(self, request, context):
        """Checks a constructed SQL query against a sql-based connection to see if it's valid based on that connection's data schema
        This is useful when constructing subsets to see if the WHERE clause is correct
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_ConnectionServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'GetConnections': grpc.unary_unary_rpc_method_handler(
                    servicer.GetConnections,
                    request_deserializer=mgmt_dot_v1alpha1_dot_connection__pb2.GetConnectionsRequest.FromString,
                    response_serializer=mgmt_dot_v1alpha1_dot_connection__pb2.GetConnectionsResponse.SerializeToString,
            ),
            'GetConnection': grpc.unary_unary_rpc_method_handler(
                    servicer.GetConnection,
                    request_deserializer=mgmt_dot_v1alpha1_dot_connection__pb2.GetConnectionRequest.FromString,
                    response_serializer=mgmt_dot_v1alpha1_dot_connection__pb2.GetConnectionResponse.SerializeToString,
            ),
            'CreateConnection': grpc.unary_unary_rpc_method_handler(
                    servicer.CreateConnection,
                    request_deserializer=mgmt_dot_v1alpha1_dot_connection__pb2.CreateConnectionRequest.FromString,
                    response_serializer=mgmt_dot_v1alpha1_dot_connection__pb2.CreateConnectionResponse.SerializeToString,
            ),
            'UpdateConnection': grpc.unary_unary_rpc_method_handler(
                    servicer.UpdateConnection,
                    request_deserializer=mgmt_dot_v1alpha1_dot_connection__pb2.UpdateConnectionRequest.FromString,
                    response_serializer=mgmt_dot_v1alpha1_dot_connection__pb2.UpdateConnectionResponse.SerializeToString,
            ),
            'DeleteConnection': grpc.unary_unary_rpc_method_handler(
                    servicer.DeleteConnection,
                    request_deserializer=mgmt_dot_v1alpha1_dot_connection__pb2.DeleteConnectionRequest.FromString,
                    response_serializer=mgmt_dot_v1alpha1_dot_connection__pb2.DeleteConnectionResponse.SerializeToString,
            ),
            'IsConnectionNameAvailable': grpc.unary_unary_rpc_method_handler(
                    servicer.IsConnectionNameAvailable,
                    request_deserializer=mgmt_dot_v1alpha1_dot_connection__pb2.IsConnectionNameAvailableRequest.FromString,
                    response_serializer=mgmt_dot_v1alpha1_dot_connection__pb2.IsConnectionNameAvailableResponse.SerializeToString,
            ),
            'CheckConnectionConfig': grpc.unary_unary_rpc_method_handler(
                    servicer.CheckConnectionConfig,
                    request_deserializer=mgmt_dot_v1alpha1_dot_connection__pb2.CheckConnectionConfigRequest.FromString,
                    response_serializer=mgmt_dot_v1alpha1_dot_connection__pb2.CheckConnectionConfigResponse.SerializeToString,
            ),
            'CheckConnectionConfigById': grpc.unary_unary_rpc_method_handler(
                    servicer.CheckConnectionConfigById,
                    request_deserializer=mgmt_dot_v1alpha1_dot_connection__pb2.CheckConnectionConfigByIdRequest.FromString,
                    response_serializer=mgmt_dot_v1alpha1_dot_connection__pb2.CheckConnectionConfigByIdResponse.SerializeToString,
            ),
            'CheckSqlQuery': grpc.unary_unary_rpc_method_handler(
                    servicer.CheckSqlQuery,
                    request_deserializer=mgmt_dot_v1alpha1_dot_connection__pb2.CheckSqlQueryRequest.FromString,
                    response_serializer=mgmt_dot_v1alpha1_dot_connection__pb2.CheckSqlQueryResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'mgmt.v1alpha1.ConnectionService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))
    server.add_registered_method_handlers('mgmt.v1alpha1.ConnectionService', rpc_method_handlers)


 # This class is part of an EXPERIMENTAL API.
class ConnectionService(object):
    """Service for managing datasource connections.
    This is a primary data model in Neosync and is used in reference when hooking up Jobs to synchronize and generate data.
    """

    @staticmethod
    def GetConnections(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/mgmt.v1alpha1.ConnectionService/GetConnections',
            mgmt_dot_v1alpha1_dot_connection__pb2.GetConnectionsRequest.SerializeToString,
            mgmt_dot_v1alpha1_dot_connection__pb2.GetConnectionsResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def GetConnection(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/mgmt.v1alpha1.ConnectionService/GetConnection',
            mgmt_dot_v1alpha1_dot_connection__pb2.GetConnectionRequest.SerializeToString,
            mgmt_dot_v1alpha1_dot_connection__pb2.GetConnectionResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def CreateConnection(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/mgmt.v1alpha1.ConnectionService/CreateConnection',
            mgmt_dot_v1alpha1_dot_connection__pb2.CreateConnectionRequest.SerializeToString,
            mgmt_dot_v1alpha1_dot_connection__pb2.CreateConnectionResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def UpdateConnection(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/mgmt.v1alpha1.ConnectionService/UpdateConnection',
            mgmt_dot_v1alpha1_dot_connection__pb2.UpdateConnectionRequest.SerializeToString,
            mgmt_dot_v1alpha1_dot_connection__pb2.UpdateConnectionResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def DeleteConnection(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/mgmt.v1alpha1.ConnectionService/DeleteConnection',
            mgmt_dot_v1alpha1_dot_connection__pb2.DeleteConnectionRequest.SerializeToString,
            mgmt_dot_v1alpha1_dot_connection__pb2.DeleteConnectionResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def IsConnectionNameAvailable(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/mgmt.v1alpha1.ConnectionService/IsConnectionNameAvailable',
            mgmt_dot_v1alpha1_dot_connection__pb2.IsConnectionNameAvailableRequest.SerializeToString,
            mgmt_dot_v1alpha1_dot_connection__pb2.IsConnectionNameAvailableResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def CheckConnectionConfig(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/mgmt.v1alpha1.ConnectionService/CheckConnectionConfig',
            mgmt_dot_v1alpha1_dot_connection__pb2.CheckConnectionConfigRequest.SerializeToString,
            mgmt_dot_v1alpha1_dot_connection__pb2.CheckConnectionConfigResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def CheckConnectionConfigById(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/mgmt.v1alpha1.ConnectionService/CheckConnectionConfigById',
            mgmt_dot_v1alpha1_dot_connection__pb2.CheckConnectionConfigByIdRequest.SerializeToString,
            mgmt_dot_v1alpha1_dot_connection__pb2.CheckConnectionConfigByIdResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def CheckSqlQuery(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/mgmt.v1alpha1.ConnectionService/CheckSqlQuery',
            mgmt_dot_v1alpha1_dot_connection__pb2.CheckSqlQueryRequest.SerializeToString,
            mgmt_dot_v1alpha1_dot_connection__pb2.CheckSqlQueryResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)
