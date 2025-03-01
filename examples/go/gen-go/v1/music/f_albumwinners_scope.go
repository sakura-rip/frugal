// Autogenerated by Frugal Compiler (3.14.2)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package music

import (
	"context"
	"fmt"

	frugal "github.com/Workiva/frugal/lib/go"
	"github.com/apache/thrift/lib/go/thrift"
)

// Scopes are a Frugal extension to the IDL for declaring PubSub
// semantics. Subscribers to this scope will be notified if they win a contest.
// Scopes must have a prefix.
type AlbumWinnersPublisher interface {
	Open() error
	Close() error
	PublishContestStart(fctx frugal.FContext, req []*Album) error
	PublishTimeLeft(fctx frugal.FContext, req Minutes) error
	PublishWinner(fctx frugal.FContext, req *Album) error
}

type albumWinnersPublisher struct {
	client  frugal.FClient
	methods map[string]*frugal.Method
}

func NewAlbumWinnersPublisher(provider *frugal.FScopeProvider, middleware ...frugal.ServiceMiddleware) AlbumWinnersPublisher {
	publisher := &albumWinnersPublisher{
		client:  frugal.NewFScopeClient(provider),
		methods: make(map[string]*frugal.Method),
	}
	middleware = append(middleware, provider.GetMiddleware()...)
	publisher.methods["publishContestStart"] = frugal.NewMethod(publisher, publisher.publishContestStart, "publishContestStart", middleware)
	publisher.methods["publishTimeLeft"] = frugal.NewMethod(publisher, publisher.publishTimeLeft, "publishTimeLeft", middleware)
	publisher.methods["publishWinner"] = frugal.NewMethod(publisher, publisher.publishWinner, "publishWinner", middleware)
	return publisher
}

func (p albumWinnersPublisher) Open() error  { return p.client.Open() }
func (p albumWinnersPublisher) Close() error { return p.client.Close() }

func (p *albumWinnersPublisher) PublishContestStart(fctx frugal.FContext, req []*Album) error {
	ret := p.methods["publishContestStart"].Invoke([]interface{}{fctx, req})
	if ret[0] != nil {
		return ret[0].(error)
	}
	return nil
}

func (p *albumWinnersPublisher) publishContestStart(fctx frugal.FContext, req []*Album) error {
	prefix := "v1.music."
	op := "ContestStart"
	topic := fmt.Sprintf("%sAlbumWinners.%s", prefix, op)
	return p.client.Publish(fctx, op, topic, albumWinnersContestStartMessage(req))
}

type albumWinnersContestStartMessage []*Album

func (p albumWinnersContestStartMessage) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteListBegin(ctx, thrift.STRUCT, len(p)); err != nil {
		return thrift.PrependError("error writing list begin: ", err)
	}
	for _, v := range p {
		if err := v.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
		}
	}
	if err := oprot.WriteListEnd(ctx); err != nil {
		return thrift.PrependError("error writing list end: ", err)
	}
	return nil
}

func (p albumWinnersContestStartMessage) Read(ctx context.Context, iprot thrift.TProtocol) error {
	panic("Not Implemented!")
}

func (p *albumWinnersPublisher) PublishTimeLeft(fctx frugal.FContext, req Minutes) error {
	ret := p.methods["publishTimeLeft"].Invoke([]interface{}{fctx, req})
	if ret[0] != nil {
		return ret[0].(error)
	}
	return nil
}

func (p *albumWinnersPublisher) publishTimeLeft(fctx frugal.FContext, req Minutes) error {
	prefix := "v1.music."
	op := "TimeLeft"
	topic := fmt.Sprintf("%sAlbumWinners.%s", prefix, op)
	return p.client.Publish(fctx, op, topic, albumWinnersTimeLeftMessage(req))
}

type albumWinnersTimeLeftMessage Minutes

func (p albumWinnersTimeLeftMessage) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteDouble(ctx, float64(p)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
	}
	return nil
}

func (p albumWinnersTimeLeftMessage) Read(ctx context.Context, iprot thrift.TProtocol) error {
	panic("Not Implemented!")
}

func (p *albumWinnersPublisher) PublishWinner(fctx frugal.FContext, req *Album) error {
	ret := p.methods["publishWinner"].Invoke([]interface{}{fctx, req})
	if ret[0] != nil {
		return ret[0].(error)
	}
	return nil
}

func (p *albumWinnersPublisher) publishWinner(fctx frugal.FContext, req *Album) error {
	prefix := "v1.music."
	op := "Winner"
	topic := fmt.Sprintf("%sAlbumWinners.%s", prefix, op)
	return p.client.Publish(fctx, op, topic, req)
}

// Scopes are a Frugal extension to the IDL for declaring PubSub
// semantics. Subscribers to this scope will be notified if they win a contest.
// Scopes must have a prefix.
type AlbumWinnersSubscriber interface {
	SubscribeContestStart(handler func(frugal.FContext, []*Album)) (*frugal.FSubscription, error)
	SubscribeTimeLeft(handler func(frugal.FContext, Minutes)) (*frugal.FSubscription, error)
	SubscribeWinner(handler func(frugal.FContext, *Album)) (*frugal.FSubscription, error)
}

// Scopes are a Frugal extension to the IDL for declaring PubSub
// semantics. Subscribers to this scope will be notified if they win a contest.
// Scopes must have a prefix.
type AlbumWinnersErrorableSubscriber interface {
	SubscribeContestStartErrorable(handler func(frugal.FContext, []*Album) error) (*frugal.FSubscription, error)
	SubscribeTimeLeftErrorable(handler func(frugal.FContext, Minutes) error) (*frugal.FSubscription, error)
	SubscribeWinnerErrorable(handler func(frugal.FContext, *Album) error) (*frugal.FSubscription, error)
}

type albumWinnersSubscriber struct {
	provider   *frugal.FScopeProvider
	middleware []frugal.ServiceMiddleware
}

func NewAlbumWinnersSubscriber(provider *frugal.FScopeProvider, middleware ...frugal.ServiceMiddleware) AlbumWinnersSubscriber {
	middleware = append(middleware, provider.GetMiddleware()...)
	return &albumWinnersSubscriber{provider: provider, middleware: middleware}
}

func NewAlbumWinnersErrorableSubscriber(provider *frugal.FScopeProvider, middleware ...frugal.ServiceMiddleware) AlbumWinnersErrorableSubscriber {
	middleware = append(middleware, provider.GetMiddleware()...)
	return &albumWinnersSubscriber{provider: provider, middleware: middleware}
}

func (l *albumWinnersSubscriber) SubscribeContestStart(handler func(frugal.FContext, []*Album)) (*frugal.FSubscription, error) {
	return l.SubscribeContestStartErrorable(func(fctx frugal.FContext, arg []*Album) error {
		handler(fctx, arg)
		return nil
	})
}

func (l *albumWinnersSubscriber) SubscribeContestStartErrorable(handler func(frugal.FContext, []*Album) error) (*frugal.FSubscription, error) {
	op := "ContestStart"
	prefix := "v1.music."
	topic := fmt.Sprintf("%sAlbumWinners.%s", prefix, op)
	transport, protocolFactory := l.provider.NewSubscriber()
	cb := l.recvContestStart(op, protocolFactory, handler)
	if err := transport.Subscribe(topic, cb); err != nil {
		return nil, err
	}

	sub := frugal.NewFSubscription(topic, transport)
	return sub, nil
}

func (l *albumWinnersSubscriber) recvContestStart(op string, pf *frugal.FProtocolFactory, handler func(frugal.FContext, []*Album) error) frugal.FAsyncCallback {
	method := frugal.NewMethod(l, handler, "SubscribeContestStart", l.middleware)
	return func(transport thrift.TTransport) error {
		iprot := pf.GetProtocol(transport)
		fctx, err := iprot.ReadRequestHeader()
		if err != nil {
			return err
		}

		ctx, cancelFn := frugal.ToContext(fctx)
		defer cancelFn()

		name, _, _, err := iprot.ReadMessageBegin(ctx)
		if err != nil {
			return err
		}

		if name != op {
			iprot.Skip(ctx, thrift.STRUCT)
			iprot.ReadMessageEnd(ctx)
			return thrift.NewTApplicationException(frugal.APPLICATION_EXCEPTION_UNKNOWN_METHOD, "Unknown function"+name)
		}
		_, size, err := iprot.ReadListBegin(ctx)
		if err != nil {
			return thrift.PrependError("error reading list begin: ", err)
		}
		req := make([]*Album, 0, size)
		for i := 0; i < size; i++ {
			elem1 := NewAlbum()
			if err := elem1.Read(ctx, iprot); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", elem1), err)
			}
			req = append(req, elem1)
		}
		if err := iprot.ReadListEnd(ctx); err != nil {
			return thrift.PrependError("error reading list end: ", err)
		}
		iprot.ReadMessageEnd(ctx)

		return method.Invoke([]interface{}{fctx, req}).Error()
	}
}

func (l *albumWinnersSubscriber) SubscribeTimeLeft(handler func(frugal.FContext, Minutes)) (*frugal.FSubscription, error) {
	return l.SubscribeTimeLeftErrorable(func(fctx frugal.FContext, arg Minutes) error {
		handler(fctx, arg)
		return nil
	})
}

func (l *albumWinnersSubscriber) SubscribeTimeLeftErrorable(handler func(frugal.FContext, Minutes) error) (*frugal.FSubscription, error) {
	op := "TimeLeft"
	prefix := "v1.music."
	topic := fmt.Sprintf("%sAlbumWinners.%s", prefix, op)
	transport, protocolFactory := l.provider.NewSubscriber()
	cb := l.recvTimeLeft(op, protocolFactory, handler)
	if err := transport.Subscribe(topic, cb); err != nil {
		return nil, err
	}

	sub := frugal.NewFSubscription(topic, transport)
	return sub, nil
}

func (l *albumWinnersSubscriber) recvTimeLeft(op string, pf *frugal.FProtocolFactory, handler func(frugal.FContext, Minutes) error) frugal.FAsyncCallback {
	method := frugal.NewMethod(l, handler, "SubscribeTimeLeft", l.middleware)
	return func(transport thrift.TTransport) error {
		iprot := pf.GetProtocol(transport)
		fctx, err := iprot.ReadRequestHeader()
		if err != nil {
			return err
		}

		ctx, cancelFn := frugal.ToContext(fctx)
		defer cancelFn()

		name, _, _, err := iprot.ReadMessageBegin(ctx)
		if err != nil {
			return err
		}

		if name != op {
			iprot.Skip(ctx, thrift.STRUCT)
			iprot.ReadMessageEnd(ctx)
			return thrift.NewTApplicationException(frugal.APPLICATION_EXCEPTION_UNKNOWN_METHOD, "Unknown function"+name)
		}
		var req Minutes
		if v, err := iprot.ReadDouble(ctx); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			temp := Minutes(v)
			req = temp
		}
		iprot.ReadMessageEnd(ctx)

		return method.Invoke([]interface{}{fctx, req}).Error()
	}
}

func (l *albumWinnersSubscriber) SubscribeWinner(handler func(frugal.FContext, *Album)) (*frugal.FSubscription, error) {
	return l.SubscribeWinnerErrorable(func(fctx frugal.FContext, arg *Album) error {
		handler(fctx, arg)
		return nil
	})
}

func (l *albumWinnersSubscriber) SubscribeWinnerErrorable(handler func(frugal.FContext, *Album) error) (*frugal.FSubscription, error) {
	op := "Winner"
	prefix := "v1.music."
	topic := fmt.Sprintf("%sAlbumWinners.%s", prefix, op)
	transport, protocolFactory := l.provider.NewSubscriber()
	cb := l.recvWinner(op, protocolFactory, handler)
	if err := transport.Subscribe(topic, cb); err != nil {
		return nil, err
	}

	sub := frugal.NewFSubscription(topic, transport)
	return sub, nil
}

func (l *albumWinnersSubscriber) recvWinner(op string, pf *frugal.FProtocolFactory, handler func(frugal.FContext, *Album) error) frugal.FAsyncCallback {
	method := frugal.NewMethod(l, handler, "SubscribeWinner", l.middleware)
	return func(transport thrift.TTransport) error {
		iprot := pf.GetProtocol(transport)
		fctx, err := iprot.ReadRequestHeader()
		if err != nil {
			return err
		}

		ctx, cancelFn := frugal.ToContext(fctx)
		defer cancelFn()

		name, _, _, err := iprot.ReadMessageBegin(ctx)
		if err != nil {
			return err
		}

		if name != op {
			iprot.Skip(ctx, thrift.STRUCT)
			iprot.ReadMessageEnd(ctx)
			return thrift.NewTApplicationException(frugal.APPLICATION_EXCEPTION_UNKNOWN_METHOD, "Unknown function"+name)
		}
		req := NewAlbum()
		if err := req.Read(ctx, iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", req), err)
		}
		iprot.ReadMessageEnd(ctx)

		return method.Invoke([]interface{}{fctx, req}).Error()
	}
}
